export PATH="/usr/local/bin:$PATH"

# Path to the bash it configuration
export BASH_IT="/Users/pivotal/.bash_it"

# Lock and Load a custom theme file
export BASH_IT_THEME="demula"
export SCM_CHECK=false
git config --global --add bash-it.hide-status 0

# Load Bash It
source $BASH_IT/bash_it.sh
