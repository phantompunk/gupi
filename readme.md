# Gupi
Gupi is a CLI to manage and render templates.

## Installation
```bash
> go install https://github.com/phantompunk/gupi@latest
```

## Usage

### Create a template
```bash
> gupi create mytemplate --template sample
```

### Render a new file
```bash
> gupi new mynewfile --template mytemplate
```

### Edit a template
```bash
> gupi edit mytemplate
```

### List all templates
```bash
> gupi list
```

### Delete a template
```bash
> gupi delete mytemplate
```

### Use the sample template

```bash
> gupi create sample --sample
```

### Use template from the web

```bash
> gupi create web -f "https://gist.githubusercontent.com/phantompunk/a3368b75e1b0ea843d12d96b949581b6/raw/c77344716ab7a69387b742fb098cec661c2ee4d7/weekly-template.md"
```

### Create a template with a dynamic name

```bash
> gupi new (date +"%Y-Week-%V.md") --template weekly --output ~/Notes/
```

## Supported Template Functions & Variables

- Date related variables:
  - `.Week`
  - `.Year`
  - `.Mon` - `.Sun`
