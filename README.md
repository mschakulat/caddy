# Caddy
## A JavaScript tool manager

### Use Caddy to install and manage your JavaScript tools
Currently supported tools:
- [Pnpm](https://pnpm.io/)
- [Node.js](https://nodejs.org/en/)

### Install Caddy
```bash
curl -s https://raw.githubusercontent.com/mschakulat/caddy/main/ci/install.sh | bash
```

### Install a tool
```bash
caddy install pnpm
caddy install node
```

### Install a specific version of a tool
```bash
caddy install pnpm@8.10.0
caddy install node@21.2.0
```

### Use a tool
```bash
pnpm -v
node -v
```

#### **If you install node with caddy, it will automatically add `npm` and `npx` to your path.**

### Use Caddy in a project
If you run Caddy in a directory with a `package.json` file,
it will automatically install the needed version of node and pnpm.
Therefor you hava to add the following to your `package.json` file:
```json
{
    "caddy": {
        "pnpm": "8.10.0",
        "node": "21.2.0"
    }
}
```

### Pin a tool to a specific version
You also can use the `caddy pin` command to add a specific version of a tool to your `package.json` file:
```bash
caddy pin pnpm@8.10.0
caddy pin node@21.2.0
```

### Uninstall Caddy
If you want to uninstall Caddy, you can simply delete the Caddy directory:
```bash
rm ~/.caddy
```
After that you have to remove the following lines from your `~/.bashrc` or `~/.zshrc` file:
```bash
export CADDY_HOME=~/.caddy
export PATH=$CADDY_HOME/bin:$PATH
```