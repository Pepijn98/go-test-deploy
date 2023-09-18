# go-test-deploy
test

This was a pain to test today
---
git clone using ssh fucked things up. Found out by cloning using https and everything seemed to work. \
This `exec bash \-l` thing might fix that, have to check tomorrow. \
Apperently it doesn't load `.bashrc` which has the keychain cmd in it, might be why ssh git clone doesn't work. \

- Change back to ssh git clone
- Test using `exec bash \-l`
- If don't work just use https git clone
