# 1. Adım: Çalışma ortamını belirle
FROM golang:1.22.0

# 2. Adım: Uygulama dosyalarını Docker içine kopyala
WORKDIR /usr/src/app
COPY . .

RUN go mod tidy
# # 3. Adım: Uygulamayı derle
# RUN go build -o main .

# # 4. Adım: Çalışma zamanı için hedef belirle
# CMD ["./main"]
