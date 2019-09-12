#!/bin/bash
echo "Regenerate all mocks..."
mockery -name Interface -dir "./pkg/storage" -output "./pkg/storage/mocks";
mockery -name Interface -dir "./pkg/twitterclient" -output "./pkg/twitterclient/mocks";