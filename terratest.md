1. Start a test
    aws-vault exec sand -- go test eks_managed_worker_test.go -v -timeout 40m 