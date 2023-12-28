<p align="center">
    <img alt="Caddy" src="./caddy.png?raw=true" width="360">
</p>

<p align="center">
    A JavaScript tool manager
</p>

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

#### **If you install node with Caddy, it will automatically add `npm` and `npx` to your path.**

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

## No conflict mode

You are free to change the section term which you are using in the `package.json`. By default Caddy uses the term `caddy`:

```json
{
    "caddy": {
        "pnpm": "8.10.0",
        "node": "21.2.0"
    }
}
```
But you can change it to whatever you want:

```bash
caddy config id <your-term>
```

After that, pinning or reading a version will respect your configured term:

```json
{
    "<your-term>": {
        "pnpm": "8.10.0",
        "node": "21.2.0"
    }
}
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
