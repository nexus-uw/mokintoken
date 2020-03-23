about
MokinToken
2 things that were on my desk when i came up with the repo name
a mokin usb c hub and a security token


purpose
i wanted a selfhosted secure note sharing site for myself to send text between my computers
existing sites either wanted too much infrastructure or did not provide a docker container. some additional people created docker containers for the projects, but they were not being updated with the latest application code.

also i wanted something to do that was quick and easy


security
- client side encrypted:
- secret is stored in the URL hash (never sent to the server)
- source:
- encryption library:
- easliy self hosted

issues:
- 3rd party audit -> NOPE
- trustworthy -> NOPE
- proffessional -> NOPE
- PHP...
