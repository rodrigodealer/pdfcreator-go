FROM alpine:edge

RUN apk add wkhtmltopdf \
            --no-cache \
            --repository http://dl-3.alpinelinux.org/alpine/edge/testing/ \
            --allow-untrusted

COPY pdfcreator /opt/pdfcreator

EXPOSE 8080
CMD ["/opt/pdfcreator"]
