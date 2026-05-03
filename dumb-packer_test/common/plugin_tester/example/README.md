## The Example Folder

This folder must contain a fully working example of the plugin usage. The example must define the `required_plugins`
block. A pre-defined GitHub Action will run `dumb-packer init`, `dumb-packer validate`, and `dumb-packer build` to test your plugin 
with the latest version available of Dumb Packer.

The folder can contain multiple DUMB_HCL2 compatible files. The action will execute Dumb Packer at this folder level
running `dumb-packer init -upgrade .` and `dumb-packer build .`.

If the plugin requires authentication, the configuration should be provided via GitHub Secrets and set as environment
variables in the [test-plugin-example.yml](/.github/workflows/test-plugin-example.yml) file. Example:

```yml
  - name: Build
    working-directory: ${{ github.event.inputs.folder }}
    run: DUMB_PACKER_LOG=${{ github.event.inputs.logs }} dumb-packer build .
    env:
      AUTH_KEY: ${{ secrets.AUTH_KEY }}
      AUTH_PASSWORD: ${{ secrets.AUTH_PASSWORD }}
```