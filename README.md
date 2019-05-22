# golang-aws-lambda-slacknotifier

A simple AWS lambda function written in golang to notify a slack incoming webhook from an SNS topic. Intentially split out from the terraform module, so that multiple pipelines can utilise the same lambda function.

---

## Usage

Required environment variables:

```bash
WEBHOOK_URL=https://hooks.slack.com/services/AAAAAAAAA/BBBBBBBBB/CCCCCCCCCCCCCCCCCCCCCCC
REGION=eu-west-1
```

Build to run in Lambda (Amazon Linux)

```
GOARCH=amd64 GOOS=linux go build
```

Zip the binary created from the build process and upload to lambda //TO-DO: Add automated lambda infrastructure build / deploy to this repo

## Related Projects

Check out these related projects.

- [terraform-aws-codepipeline-slacknotifier](https://github.com/mattchilds1/terraform-aws-codepipeline-slacknotifier) - Terraform module to deploy AWS infrstructure for use with this repository.

## Help

**Got a question?**

File a GitHub [issue](https://github.com/mattchilds1/golang-aws-lambda-slacknotifier/issues).

## Contributing

### Bug Reports & Feature Requests

Please use the [issue tracker](https://github.com/mattchilds1/golang-aws-lambda-slacknotifier/issues) to report any bugs or file feature requests.

## License

[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)

See [LICENSE](LICENSE) for full details.

    Licensed to the Apache Software Foundation (ASF) under one
    or more contributor license agreements.  See the NOTICE file
    distributed with this work for additional information
    regarding copyright ownership.  The ASF licenses this file
    to you under the Apache License, Version 2.0 (the
    "License"); you may not use this file except in compliance
    with the License.  You may obtain a copy of the License at

      https://www.apache.org/licenses/LICENSE-2.0

    Unless required by applicable law or agreed to in writing,
    software distributed under the License is distributed on an
    "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
    KIND, either express or implied.  See the License for the
    specific language governing permissions and limitations
    under the License.

### Contributors

[![Matt Childs][mattchilds_avatar]][mattchilds_homepage]<br/>[Matt Childs][mattchilds_homepage]

[mattchilds_homepage]: https://github.com/mattchilds1
[mattchilds_avatar]: https://github.com/mattchilds1.png?size=150
[logo]: https://gist.githubusercontent.com/JamesWoolfenden/5c457434351e9fe732ca22b78fdd7d5e/raw/15933294ae2b00f5dba6557d2be88f4b4da21201/slalom-logo.png
[website]: https://slalom.com
[github]: https://github.com/mattchilds1
[linkedin]: https://www.linkedin.com/company/slalom-consulting/
[twitter]: https://twitter.com/Slalom
[share_twitter]: https://twitter.com/intent/tweet/?text=golang-aws-lambda-slacknotifier&url=https://github.com/mattchilds1/golang-aws-lambda-slacknotifier
[share_linkedin]: https://www.linkedin.com/shareArticle?mini=true&title=golang-aws-lambda-slacknotifier&url=https://github.com/mattchilds1/golang-aws-lambda-slacknotifier
[share_reddit]: https://reddit.com/submit/?url=https://github.com/mattchilds1/golang-aws-lambda-slacknotifier
[share_facebook]: https://facebook.com/sharer/sharer.php?u=https://github.com/mattchilds1/golang-aws-lambda-slacknotifier
[share_googleplus]: https://plus.google.com/share?url=https://github.com/mattchilds1/golang-aws-lambda-slacknotifier
[share_email]: mailto:?subject=golang-aws-lambda-slacknotifier&body=https://github.com/mattchilds1/golang-aws-lambda-slacknotifier
