resource "aws_lambda_function" "slack_notifier" {
  function_name    = "${var.function_name}"
  filename         = "../build.zip"
  handler          = "slackNotifier"
  source_code_hash = "${base64sha256(file("../build.zip"))}"
  role             = "${aws_iam_role.slack_notifier.arn}"
  runtime          = "go1.x"
  memory_size      = 128
  timeout          = 1

  environment {
    variables = {
      WEBHOOK_URL = "${var.slack_webhook}"
      REGION      = "${var.region}"
    }
  }
}
