-- Neovim project-specific Lua configuration for blockfont

-- Set Go-specific options
vim.opt_local.tabstop = 4
vim.opt_local.shiftwidth = 4
vim.opt_local.expandtab = false

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

-- Format on save for Go files
vim.api.nvim_create_autocmd("BufWritePre", {
  pattern = "*.go",
  callback = function()
    vim.lsp.buf.format({ async = false })
  end,
})

-- Run tests on file save
vim.api.nvim_create_autocmd("BufWritePost", {
  pattern = "*_test.go",
  callback = function()
    vim.notify("Running tests...", vim.log.levels.INFO)
    vim.fn.jobstart("go test -v ./...", {
      on_exit = function(_, code)
        if code == 0 then
          vim.notify("Tests passed!", vim.log.levels.INFO)
        else
          vim.notify("Tests failed!", vim.log.levels.ERROR)
        end
      end,
    })
  end,
})
