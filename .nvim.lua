-- blockfont project-specific Neovim configuration
-- This file is automatically loaded when opening Neovim in the project directory

-- Set Go-specific options
vim.opt_local.tabstop = 4
vim.opt_local.shiftwidth = 4
vim.opt_local.expandtab = false

-- Project root detection
vim.g.blockfont_root = vim.fn.getcwd()

-- Custom commands for this project
vim.api.nvim_create_user_command('BlockfontBuild', '!make build', {})
vim.api.nvim_create_user_command('BlockfontTest', '!make test', {})
vim.api.nvim_create_user_command('BlockfontDemo', '!make demo', {})
vim.api.nvim_create_user_command('BlockfontFmt', '!make fmt', {})
vim.api.nvim_create_user_command('BlockfontLint', '!make lint', {})
vim.api.nvim_create_user_command('BlockfontDocsDev', '!make docs-dev &', {})
vim.api.nvim_create_user_command('BlockfontDocsBuild', '!make docs-build', {})
vim.api.nvim_create_user_command('BlockfontDocsOpen', '!make docs-open', {})
vim.api.nvim_create_user_command('BlockfontDemoRecord', '!nix run .#demo-record', {})
vim.api.nvim_create_user_command('BlockfontDemoPlay', '!nix run .#demo-play', {})

-- Which-key integration for project keybindings (new spec format)
local ok, wk = pcall(require, "which-key")
if ok then
  wk.add({
    { "<leader>p", group = "Project (blockfont)" },

    -- Build & Run
    { "<leader>pb", "<cmd>!make build<cr>", desc = "Build library" },
    { "<leader>pd", "<cmd>!make demo<cr>", desc = "Run demo application" },
    { "<leader>pt", "<cmd>!make test<cr>", desc = "Run tests" },
    { "<leader>pT", "<cmd>!make test-coverage<cr>", desc = "Run tests with coverage" },

    -- Examples
    { "<leader>pe", group = "Examples" },
    { "<leader>pes", "<cmd>!make run-simple<cr>", desc = "Run simple example" },
    { "<leader>pea", "<cmd>!make run-animated<cr>", desc = "Run animated example" },
    { "<leader>pee", "<cmd>!make run-editor<cr>", desc = "Run editor example" },

    -- Code Quality
    { "<leader>pf", "<cmd>!make fmt<cr>", desc = "Format code" },
    { "<leader>pl", "<cmd>!make lint<cr>", desc = "Lint code" },

    -- Documentation (MkDocs)
    { "<leader>pD", group = "Documentation" },
    { "<leader>pDd", "<cmd>!make docs-dev &<cr>", desc = "Start MkDocs dev server" },
    { "<leader>pDb", "<cmd>!make docs-build<cr>", desc = "Build documentation" },
    { "<leader>pDc", "<cmd>!make docs-clean<cr>", desc = "Clean documentation" },
    { "<leader>pDo", "<cmd>!make docs-open<cr>", desc = "Open docs in browser" },

    -- Demo Recording (asciinema)
    { "<leader>pr", group = "Recording" },
    { "<leader>prr", "<cmd>terminal nix run .#demo-record<cr>", desc = "Record demo" },
    { "<leader>prp", "<cmd>terminal nix run .#demo-play<cr>", desc = "Play demo" },

    -- Release
    { "<leader>pR", "<cmd>terminal nix run .#release<cr>", desc = "Release new version" },

    -- Nix
    { "<leader>pn", group = "Nix" },
    { "<leader>pnb", "<cmd>!nix build<cr>", desc = "Nix build" },
    { "<leader>pnr", "<cmd>!nix run<cr>", desc = "Nix run (demo)" },
    { "<leader>pnd", "<cmd>!nix run .#docs-serve<cr>", desc = "Nix docs serve" },
    { "<leader>pnt", "<cmd>!nix run .#test<cr>", desc = "Nix test" },

    -- Version & Dependencies
    { "<leader>pv", "<cmd>!make version<cr>", desc = "Show version info" },
    { "<leader>pu", "<cmd>!make deps<cr>", desc = "Update dependencies (go mod tidy)" },
  })
end

-- Autocommands for this project
local blockfont_group = vim.api.nvim_create_augroup("BlockfontProject", { clear = true })

-- Format Go files on save
vim.api.nvim_create_autocmd("BufWritePre", {
  group = blockfont_group,
  pattern = "*.go",
  callback = function()
    vim.lsp.buf.format({ async = false })
  end,
})

-- Run tests on test file save (optional, can be noisy)
-- Uncomment if you want automatic test runs
-- vim.api.nvim_create_autocmd("BufWritePost", {
--   group = blockfont_group,
--   pattern = "*_test.go",
--   callback = function()
--     vim.notify("Running tests...", vim.log.levels.INFO)
--     vim.fn.jobstart("go test -v ./...", {
--       on_exit = function(_, code)
--         if code == 0 then
--           vim.notify("Tests passed!", vim.log.levels.INFO)
--         else
--           vim.notify("Tests failed!", vim.log.levels.ERROR)
--         end
--       end,
--     })
--   end,
-- })

-- LSP configuration for gopls
vim.lsp.config.gopls = {
  settings = {
    gopls = {
      analyses = {
        unusedparams = true,
        shadow = true,
      },
      staticcheck = true,
      gofumpt = true,
    },
  },
}

-- Print confirmation
vim.notify("blockfont project configuration loaded", vim.log.levels.INFO)
