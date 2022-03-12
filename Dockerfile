FROM golang:latest

ENV DATABASE_URL=postgres://postgres:docker@172.17.0.3:5432/world?sslmode=disable
