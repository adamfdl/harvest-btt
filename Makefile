export HARVEST_ACCOUNT_ID = 254107
export HARVEST_ACCESS_TOKEN = 1371640.pt.NoOXbZ6zFTKiOGsX7siUm89hpFrstiY7VPYOWrL6coug27yprzWYRrmCAI81HhgJjV1rGhpLmmVA6Gk6wFDThA

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