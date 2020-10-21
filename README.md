# Terrafrom credentials helper for Google Cloud Storage

`terraform-credentials-gcs` is a Terraform ["credentials helper"](https://www.terraform.io/docs/internals/credentials-helpers.html) plugin that allows providing credentials for Terraform-native services (private module registries, Terraform Cloud, etc) publish in private Google Cloud Storage bucket.

To use it, download a release archive and extract it into the [`~/.terraform.d/plugins`](https://www.terraform.io/docs/extend/how-terraform-works.html#plugin-locations) directory where Terraform looks for credentials helper plugins. (The filename of the file inside the archive is important for Terraform to discover it correctly, so don't rename it).

Terraform will take the newest version of the plugin it finds in the plugin search directory, so if you are switching between versions you may prefer to remove existing installed versions in order to ensure Terraform selects the desired version.

Once you've installed the plugin, enable it by adding the following block to your [Terraform CLI configuration](https://www.terraform.io/docs/commands/cli-config.html): `credentials_helper "gcs" {}`.

This credentials helper plugin does not take any additional arguments, so the block must be left empty as shown above.

To enable authentication on Google Cloud Storage, use **at least**:

- Use [`gcloud`](https://cloud.google.com/sdk/docs/install) with `gcloud auth login --update-adc --no-launch-browser`
- `export GOOGLE_APPLICATION_CREDENTIALS=/path/to/your/service/account/file.json`: file path
- `export GOOGLE_CREDENTIALS=$(cat /path/to/your/service/account/file.json | tr -d "\n")`: file content

This will provided authentication to [`golang.org/x/oauth2/google`](https://pkg.go.dev/golang.org/x/oauth2/google) wich understand these methods. See https://cloud.google.com/docs/authentication/production.

This helper will store credentials for registries in `$HOME/.config/terraform-credentials-gcs`.
