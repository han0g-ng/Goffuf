1. Phase 1: <br/>

Phiên bản ffuf cơ bản, chỉ có chức năng nhận URL và gửi đi request, sau đó lấy các thông tin phản hồi như Status Code, Content Length, Duration Time và in ra.

Sử dụng http.Client: Một đối tượng quản lý vòng đời của Request (bao gồm nhiều thành phần bên trong như timeout, redirect, cookie). <br/>
Nó là một lớp high-level, sẽ gọi xuống các lớp ở dưới như http.Transport để xử lý các thành phần trong giao thức như kết nối, bắt tay.

Cơ chế connection reuse / pooling <br/>
Một HTTP request mới phải trải qua rất nhiều bước: DNS lookup, TCP handshake, TLS handshake, gửi request. Ý tưởng là thay vì đóng request thì giữ và dùng lại. Ở đây dẫn ta đến khái niệm connection pooling: Tức http.Transport giữ cho connection rảnh và active, TCP connection không bị đóng (do không gửi FIN, kernel giữ socket mở, HTTP phải có keep-alive - tức client và server thống nhất giữ connection này còn dùng tiếp). <br/>
Phải đóng phần body để http.Transport biết được để reuse connection.


2. Phase 2 <br/>
Bắt đầu lấy wordlist, duyệt từng keyword
Trong url mục tiêu, bắt đầu sử dụng cơ chế có từ khoá FUZZ
Ta thay thế FUZZ bằng keyword trong wordlist và tiến hành truy vấn tới (vẫn sử dụng cách gửi request như trong phase 1) <br />
Lưu ý cơ chế của goroutine: Nếu sử dụng unbuffered channel, việc gửi ch <- x và <-ch (gửi và nhận) phải xảy ra đồng thời. Ta cần 2 goroutine để chạy riêng từng cái, từ đó mới không xảy ra deadlock. Việc tạo thêm go func() là để tạo 1 goroutine khác, chạy song song với goroutine cũ, trong goroutine mới này chạy send, còn cái cũ chạy receive.

3. Phase 3 <br/>
Chuyển từ xử lý tuần tự sang đồng thời. <br/>
Kiến trúc hệ thống: Mô hình Producer - Consumer<br/>
Tránh tạo ra quá nhiều luồng (goroutines), ta chỉ tận dụng số lượng có hạn và cố định. <br/>
* **Producer**: đọc wordlist theo luồng, đọc từng dòng trong file và đẩy keyword vào một kênh.
* **Job Queue**: Băng chuyền dữ liệu: dữ liệu được đưa vào 1 đầu bởi producer, và lấy ra bởi consumer
* **Worker Pool**: Đứng đợi ở đầu ra dữ liệu, các worker rảnh rỗi sẽ lấy dữ liệu đó và làm việc, tránh được quá trình nhàn rỗi. <br />
Các điểm hạn chế: <br />
* Có nhiều công nhân nhưng chỉ có 1 băng chuyền dữ liệu
* Các worker còn đang kiêm thêm nhiệm vụ in kết quả thay vì chỉ thực hiện request.
Bổ sung thêm 1 băng chuyền thứ 2 để nhận kết quả

4. Phase 4 <br/>
Bổ sung một vài filter cơ bản: chọn loại status code muốn lấy, muốn bỏ, chọn content length muốn bỏ.

5. Phase 5 <br />
Chuyển đổi sang môi trường dòng lệnh 