echo "Run gonkey tests..."
docker run -it -v "./tests:/tests" lamoda/gonkey -tests /tests/gonkey -host localhost:1308