# Webtest

This is a simple test webserver for demonstrating serveral thing with Docker and Kubernetes

## Configuration

```bash
# Envvars

# Path to the file which holds the content that gets served /filecontent
CONTENTFILE=foobar.txt

# Path to the dir which will be served under /contentdir/
CONTENTDIR=/content

# This string will be shown on /
GREETING="webtest-pod"
```