test-all:
	@go test -v ./...

test-1:
	@cd 1.reading-file && go test -v ./...

test-2:
	@cd 2.palindrome && go test -v ./...

test-3:
	@cd 3.fetch && go test -v ./...

test-4:
	@cd 4.random-number && go test -v ./...

test-5:
	@cd 5.convert-time && go test -v ./...

test-6:
	@cd 6.logic && go test -v ./...