FROM scratch
LABEL maintainer="Prasanta Singh <singh.prasanta@gmail.com>"

COPY . .
CMD ["./web-echo"]
 
EXPOSE 80 80
