Open Threat Exchange Go API Client
-----------------------------

![alt tag](https://i.imgur.com/I0USmqj.png)

### Overview

A working client implementation for https://otx.alienvault.com/ written in Golang
Currently supports:
Validate api key / obtain user statistics
Get Subscriptions
Get Pulse (with IOCs) by pulse_id

### Installing

```
$ go get -u github.com/AlienVault-Labs/OTX-Go-SDK/src/otaxpi
```

### License
See LICENSE

### Authors
Bill Smartt, Security Engineer (@bsmartt13)

### Contributing
There's likely a lot of room for improvement here by a Golang expert.  Please send PRs <3


### Issues
Please file any issues you find on github.com

### API Keys
If you haven't signed up, please visit [https://otx.alienvault.com/](https://otx.alienvault.com/accounts/signup/) and do so!

Then, you can login and find your api key in your [account settings](https://otx.alienvault.com/settings/).

Once you have your API Key, set it as an environment variable:
```export X_OTX_API_KEY=mysecretkey```
```echo $X_OTX_API_KEY```

or in your go code:
```os.Setenv("X_OTX_API_KEY", "ab91e98e6dcac6303bd1522d3542f91fcb4be176ea262ecd892d39e0d82a218b")```

For use with curl, or to write your own client, set the api key as a HTTP Header:
```X-OTX-API-KEY: "ab91e98e6dcac6303bd1522d3542f91fcb4be176ea262ecd892d39e0d82a218b"```


### Subscriptions
Your subscriptions include all pulses:
- Created by authors you subscribe to
- Pulses you subscribe to directly
- Pulses you create
