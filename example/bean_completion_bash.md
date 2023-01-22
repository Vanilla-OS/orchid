---
title: "bean completion bash"
description: bean completion bash

---
## bean completion bash

Generate the autocompletion script for bash

### Synopsis

Generate the autocompletion script for the bash shell.

This script depends on the 'bash-completion' package.
If it is not installed already, you can install it via your OS's package manager.

To load completions in your current shell session:

	source <(bean completion bash)

To load completions for every new session, execute once:

#### Linux:

	bean completion bash > /etc/bash_completion.d/bean

#### macOS:

	bean completion bash > $(brew --prefix)/etc/bash_completion.d/bean

You will need to start a new shell for this setup to take effect.


```
bean completion bash
```

### Options

```
  -h, --help              help for bash
      --no-descriptions   disable completion descriptions
```

### Options inherited from parent commands

```
  -d, --doit   do the thing
```

### SEE ALSO

* [bean completion](bean_completion/)	 - Generate the autocompletion script for the specified shell

