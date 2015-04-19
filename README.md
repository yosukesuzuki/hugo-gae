[![Circle CI](https://circleci.com/gh/yosukesuzuki/hugo-gae.svg?style=svg)](https://circleci.com/gh/yosukesuzuki/hugo-gae)

# hugo-gae
[hugo](http://gohugo.io/) and Google App Engine/Go static file serving example.

# install hugo
for mac user

```bash
$ brew install hugo
```

# install google cloud sdk
see this document
https://cloud.google.com/sdk/#Quick_Start


# deploy by circleci
generate oauth2 refresh token

```
$ python ~/google-cloud-sdk/platform/google_appengine/appcfg.py update . --oauth2
```

and then check the "refresh_token" value

```
$ cat ~/.appcfg_oauth2_tokens
```

set Environ variable in Project Settings > Environment Variables

Name is "APPENGINE_TOKEN"
Value is value of refresh_token
![Project_settings_-_yosukesuzuki_hugo-gae_-_CircleCI.png](https://qiita-image-store.s3.amazonaws.com/0/45686/f294b5b0-3045-58c5-1147-88a0300d4c17.png "Project_settings_-_yosukesuzuki_hugo-gae_-_CircleCI.png")