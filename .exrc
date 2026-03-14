" Neovim project-specific configuration for blockfont
" Load with :source .exrc or use exrc option

" Which-key mappings under <leader>p for project commands
lua << EOF
local wk = require("which-key")

wk.add({
  { "<leader>p", group = "Project" },
  { "<leader>pb", "<cmd>!go build ./...<CR>", desc = "Build" },
  { "<leader>pt", "<cmd>!go test -v ./...<CR>", desc = "Test" },
  { "<leader>pl", "<cmd>!golangci-lint run ./...<CR>", desc = "Lint" },
  { "<leader>pf", "<cmd>!gofmt -w .<CR>", desc = "Format" },
  { "<leader>pd", "<cmd>!cd docs && mkdocs serve &<CR>", desc = "Docs serve" },
  { "<leader>pr", "<cmd>!go run ./examples/simple<CR>", desc = "Run example" },
  { "<leader>pm", "<cmd>!go mod tidy<CR>", desc = "Mod tidy" },
})
EOF
