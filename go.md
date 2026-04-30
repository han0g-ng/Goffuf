1. Phase 1: <br/>

Phiên bản ffuf cơ bản, chỉ có chức năng nhận URL và gửi đi request, sau đó lấy các thông tin phản hồi như Status Code, Content Length, Duration Time và in ra.

Sử dụng http.Client: Một đối tượng quản lý vòng đời của Request (bao gồm nhiều thành phần bên trong như timeout, redirect, cookie). <br/>
Nó là một lớp high-level, sẽ gọi xuống các lớp ở dưới như http.Transport để xử lý các thành phần trong giao thức như kết nối, bắt tay.

Cơ chế connection reuse / pooling <br/>
Một HTTP request mới phải trải qua rất nhiều bước: DNS lookup, TCP handshake, TLS handshake, gửi request. Ý tưởng là thay vì đóng request thì giữ và dùng lại. Ở đây dẫn ta đến khái niệm connection pooling: Tức http.Transport giữ cho connection rảnh và active, TCP connection không bị đóng (do không gửi FIN, kernel giữ socket mở, HTTP phải có keep-alive - tức client và server thống nhất giữ connection này còn dùng tiếp). <br/>
Phải đóng phần body để http.Transport biết được để reuse connection.


2. Phase 2