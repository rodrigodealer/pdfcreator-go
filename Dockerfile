FROM alpine:3.5

RUN apk add --update --no-cache \
    libgcc libstdc++ libx11 glib libxrender libxext libintl \
    libcrypto1.0 libssl1.0 \
    ttf-dejavu ttf-droid ttf-freefont ttf-liberation ttf-ubuntu-font-family

COPY wkhtmltopdf /usr/bin/wkhtmltopdf

COPY pdfcreator* /opt/pdfcreator

CMD ["/opt/pdfcreator"]
