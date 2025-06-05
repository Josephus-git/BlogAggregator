module github.com/josephus-git/BlogAggregator

replace github.com/josephus-git/BlogAggregator/internal/config => ./internal/config

replace github.com/josephus-git/BlogAggregator/internal/database => ./internal/database

require github.com/josephus-git/BlogAggregator/internal/config v0.0.0

require github.com/josephus-git/BlogAggregator/internal/database v0.0.0

require github.com/lib/pq v1.10.9

require github.com/google/uuid v1.6.0 

go 1.22.2
