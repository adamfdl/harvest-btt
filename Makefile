export HARVEST_ACCOUNT_ID = <YOUR HARVEST ACCOUNT ID>
export HARVEST_ACCESS_TOKEN = <YOUR HARVEST ACCESS TOKEN>

test:
	go test -v -cover ./...

test-report:
	./harvest-btt reports

test-report-non-bil:
	./harvest-btt reports --non-billables

test-timer-toggle:
	./harvest-btt timer --toggle

test-timer-latest:
	./harvest-btt timer