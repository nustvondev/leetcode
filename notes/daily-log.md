# 📅 Daily Practice Log

Nhật ký luyện tập LeetCode & DSA.  
Mỗi ngày mình sẽ ghi lại: **Ngày / Bài tập / Ý tưởng / Ghi chú**.  
👉 Mục tiêu: duy trì thói quen giải 1-2 bài/ngày, ưu tiên hiểu **tư duy & trade-off** giữa các giải pháp.

---

## Week 1

### Day 01 (2025-08-26)
- **Top75**: Two Sum ✅
  - **Ý tưởng**:
    - **Brute Force**:  
      - Duyệt qua tất cả các cặp số `(i, j)`.  
      - Nếu `nums[i] + nums[j] == target` → return kết quả.  
      - **Độ phức tạp**:  
        - ⏱️ Thời gian: `O(n^2)`  
        - 💾 Không gian: `O(1)`  
    - **Hash Map (Tối ưu)**:  
      - Duyệt qua từng số, tính `complement = target - num`.  
      - Kiểm tra nếu `complement` đã có trong hash map → return chỉ mục.  
      - Nếu chưa có → thêm `num` vào hash map với chỉ mục của nó.  
      - **Độ phức tạp**:  
        - ⏱️ Thời gian: `O(n)`  
        - 💾 Không gian: `O(n)`  
  - **Ghi chú**:  
    - Hash Map là giải pháp tối ưu.  
    - Bài toán này luyện khả năng dùng hash map để kiểm tra **sự tồn tại** và **tìm kiếm nhanh**.

---

### Day 02 (2025-08-27)
- **Top75**: Contains Duplicate ✅
  - **Ý tưởng**:
    - **Brute Force**:  
      - So sánh từng cặp phần tử `(i, j)`.  
      - Nếu có `nums[i] == nums[j]` → return `true`.  
      - **Độ phức tạp**:  
        - ⏱️ Thời gian: `O(n^2)`  
        - 💾 Không gian: `O(1)`  
    - **Hash Map / Hash Set (Tối ưu)**:  
      - Dùng hash map/set để lưu các số đã gặp.  
      - Nếu gặp số đã tồn tại trong map/set → return `true`.  
      - Nếu duyệt hết không trùng → return `false`.  
      - **Độ phức tạp**:  
        - ⏱️ Thời gian: `O(n)`  
        - 💾 Không gian: `O(n)`  
    - **Sorting**:  
      - Sắp xếp mảng trước.  
      - Duyệt kiểm tra xem có phần tử nào giống phần tử trước đó.  
      - **Độ phức tạp**:  
        - ⏱️ Thời gian: `O(n log n)`  
        - 💾 Không gian: phụ thuộc vào thuật toán sort (trung bình `O(1)` đến `O(n)`).  
  - **Ghi chú**:  
    - Hash Set là giải pháp tối ưu nhất cho thời gian.  
    - Sorting hữu ích khi cần tiết kiệm bộ nhớ hoặc muốn xử lý theo thứ tự.  
    - Brute Force chỉ để tham khảo, không nên dùng thực tế.  

---

### Day 03 (2025-08-28)
- **Top75**: Best Time to Buy and Sell Stock ✅
  - **Ý tưởng**:
    - **Brute Force**:  
      - Duyệt qua mọi cặp `(buy, sell)` với `buy < sell`.  
      - Tính `profit = prices[sell] - prices[buy]`.  
      - Giữ lại lợi nhuận lớn nhất.  
      - **Độ phức tạp**:  
        - ⏱️ Thời gian: `O(n^2)`  
        - 💾 Không gian: `O(1)`  
    - **One-Pass (Tối ưu)**:  
      - Duyệt mảng 1 lần, theo dõi **minPrice** (giá nhỏ nhất từ đầu đến hiện tại).  
      - Với mỗi giá `price`, tính `profit = price - minPrice`.  
      - Cập nhật `maxProfit` nếu `profit` lớn hơn.  
      - **Độ phức tạp**:  
        - ⏱️ Thời gian: `O(n)`  
        - 💾 Không gian: `O(1)`  
  - **Ghi chú**:  
    - Đây là bài “khởi đầu” cho chuỗi bài **Stock Problems**.  
    - Tư duy quan trọng: **theo dõi min giá trước đó** và **so sánh chênh lệch**.  
    - Các biến thể nâng cao (Stock II, III, IV, k transactions) sẽ cần DP.

---

### Day 04 (2025-08-29)
- **Top75**: Valid Anagram ✅
  - **Ý tưởng**:
    - **Brute Force (Sort + So sánh)**:  
      - Sắp xếp 2 chuỗi rồi so sánh từng ký tự.  
      - Nếu giống nhau → là anagram.  
      - **Độ phức tạp**:  
        - ⏱️ Thời gian: `O(n log n)`  
        - 💾 Không gian: `O(1)` hoặc `O(n)` tùy sort.  
    - **Hash Map (Tối ưu cho lowercase)**:  
      - Đếm tần suất ký tự trong chuỗi `s`.  
      - Trừ đi tần suất ký tự trong chuỗi `t`.  
      - Nếu có ký tự nào âm → không phải anagram.  
      - **Độ phức tạp**:  
        - ⏱️ Thời gian: `O(n)`  
        - 💾 Không gian: `O(1)` (26 ký tự tiếng Anh).  
    - **Unicode-Friendly (Tổng quát)**:  
      - Sử dụng hash map để đếm tần suất ký tự dạng `rune`.  
      - Áp dụng cho mọi ngôn ngữ, kể cả Unicode.  
      - **Độ phức tạp**:  
        - ⏱️ Thời gian: `O(n)`  
        - 💾 Không gian: `O(k)` với `k` là số ký tự khác nhau.  
  - **Ghi chú**:  
    - Hash Map là cách tối ưu cho input chữ thường.  
    - Với Unicode, cần dùng map rune để xử lý tổng quát.  
    - Bài này rèn luyện kỹ năng **count frequency** – nền tảng cho nhiều bài liên quan đến chuỗi.  

---


### Day 05 (2025-09-03)
- **Top75**: Valid Parentheses ✅
  - **Ý tưởng**:
    - **Brute Force (String Replace)**:  
      - Liên tục thay thế các cặp `"()"`, `"{}"`, `"[]"` bằng chuỗi rỗng.  
      - Nếu cuối cùng chuỗi rỗng → hợp lệ.  
      - **Độ phức tạp**:  
        - ⏱️ Thời gian: `O(n^2)` trong trường hợp xấu.  
        - 💾 Không gian: `O(n)` do tạo chuỗi mới.  
    - **Stack (Chuẩn – Best Practice)**:  
      - Duyệt từng ký tự trong chuỗi:  
        - Nếu là ngoặc mở → push vào stack.  
        - Nếu là ngoặc đóng → kiểm tra top của stack có khớp không.  
      - Cuối cùng stack rỗng → hợp lệ.  
      - **Độ phức tạp**:  
        - ⏱️ Thời gian: `O(n)`  
        - 💾 Không gian: `O(n)`  
    - **Optimized Stack (Push ngoặc cần đóng)**:  
      - Khi gặp `(` thì push `)`, gặp `[` thì push `]`, gặp `{` thì push `}`.  
      - Khi gặp ngoặc đóng, chỉ cần so với top stack.  
      - **Độ phức tạp**:  
        - ⏱️ Thời gian: `O(n)`  
        - 💾 Không gian: `O(n)`  
  - **Ghi chú**:  
    - Stack là giải pháp tối ưu và dễ cài đặt nhất.  
    - Cách push “ngoặc cần đóng” giúp code gọn hơn.  
    - Đây là bài rèn luyện tư duy **stack ứng dụng cho kiểm tra cấu trúc** (parsing, compiler, expression).  


---


### Day 06 (2025-09-04)
- **Top75**: Maximum Subarray ✅
  - **Ý tưởng**:
    - **Brute Force (O(n³))**:  
      - Duyệt tất cả các subarray `(i..j)`.  
      - Tính tổng từng subarray, lấy max.  
      - **Độ phức tạp**:  
        - ⏱️ Thời gian: `O(n^3)`  
        - 💾 Không gian: `O(1)`  
    - **Prefix Sum (O(n²))**:  
      - Tính prefix sum: `prefix[i] = tổng nums[0..i-1]`.  
      - Subarray sum `(i..j) = prefix[j+1] - prefix[i]`.  
      - Giảm 1 vòng lặp so với brute force.  
      - **Độ phức tạp**:  
        - ⏱️ Thời gian: `O(n^2)`  
        - 💾 Không gian: `O(n)`  
    - **Kadane’s Algorithm (O(n)) – Best Practice**:  
      - Duyệt mảng 1 lần:  
        - Với mỗi `nums[i]`, chọn giữa bắt đầu subarray mới (`nums[i]`) hoặc nối tiếp subarray trước (`currSum + nums[i]`).  
        - Cập nhật `maxSum`.  
      - **Độ phức tạp**:  
        - ⏱️ Thời gian: `O(n)`  
        - 💾 Không gian: `O(1)`  
    - **Divide & Conquer (O(n log n))**:  
      - Chia đôi mảng thành left và right.  
      - Kết quả tối đa là max của:  
        1. Subarray max bên trái  
        2. Subarray max bên phải  
        3. Subarray max đi qua mid  
      - Dùng đệ quy để giải từng phần.  
      - **Độ phức tạp**:  
        - ⏱️ Thời gian: `O(n log n)`  
        - 💾 Không gian: `O(log n)` (stack đệ quy)  
  - **Ghi chú**:  
    - Kadane’s là chuẩn nhất (O(n), O(1)).  
    - Brute force & prefix sum giúp hiểu bản chất nhưng không thực tế.  
    - Divide & Conquer hữu ích trong phỏng vấn để thể hiện tư duy thuật toán.  
    - Đây là bài cực kỳ kinh điển, thường được hỏi trong interview để kiểm tra khả năng tối ưu tư duy.  


---


### Day 07 (2025-09-05)
- **Top75**: Product of Array Except Self ✅
  - **Ý tưởng**:
    - **Brute Force (O(n²))**:  
      - Với mỗi `i`, tính tích tất cả phần tử trừ `nums[i]`.  
      - **Độ phức tạp**:  
        - ⏱️ Thời gian: `O(n²)`  
        - 💾 Không gian: `O(1)`  
    - **Prefix & Suffix Arrays (O(n), O(n))**:  
      - Lưu `prefix[i] = tích từ trái`.  
      - Lưu `suffix[i] = tích từ phải`.  
      - Kết quả = `prefix[i] * suffix[i]`.  
      - **Độ phức tạp**:  
        - ⏱️ Thời gian: `O(n)`  
        - 💾 Không gian: `O(n)`  
    - **Optimized Prefix-Suffix (O(n), O(1)) – Best Practice**:  
      - Dùng `res[]` để lưu prefix.  
      - Một vòng từ phải sang trái để nhân suffix.  
      - **Độ phức tạp**:  
        - ⏱️ Thời gian: `O(n)`  
        - 💾 Không gian: `O(1)` (không tính `res`).  
  - **Ghi chú**:  
    - Đây là dạng bài kinh điển cho kỹ thuật **prefix-suffix**.  
    - Hạn chế “không dùng chia” giúp luyện tư duy tối ưu.  
    - Phiên bản O(1) space là chuẩn để nhớ và dùng trong interview.  

---


### Day 08 (2025-09-08)
- **Top75**: 3Sum ✅
  - **Ý tưởng**:
    - **Brute Force (O(n³))**:  
      - Duyệt tất cả bộ ba `(i, j, k)`.  
      - Nếu tổng = 0 thì lưu kết quả (dùng set để tránh trùng).  
      - **Độ phức tạp**:  
        - ⏱️ Thời gian: `O(n³)`  
        - 💾 Không gian: `O(n²)` để lưu kết quả.  
    - **Sorting + Two Pointers (O(n²)) – Best Practice**:  
      - Sort mảng trước.  
      - Với mỗi `nums[i]`, dùng hai con trỏ `l` và `r` để tìm `nums[l] + nums[r] = -nums[i]`.  
      - Nếu sum = 0 → lưu triplet, di chuyển con trỏ và skip duplicates.  
      - Nếu sum < 0 → tăng `l`.  
      - Nếu sum > 0 → giảm `r`.  
      - **Độ phức tạp**:  
        - ⏱️ Thời gian: `O(n²)`  
        - 💾 Không gian: `O(1)` (không tính output).  
  - **Ghi chú**:  
    - Đây là bài kinh điển dùng **Two Pointers + Sorting**.  
    - Cần cẩn thận xử lý **duplicate triplets**.  
    - Là nền tảng cho các biến thể nâng cao: **4Sum**, **k-Sum**, …  

---

### Day 09 (2025-09-09)
- **Top75**: Merge Intervals ✅
  - **Ý tưởng**:
    - **Brute Force (O(n²))**:  
      - So sánh từng cặp interval để xem có overlap không.  
      - Nếu có → merge, lặp lại cho đến khi không merge được nữa.  
      - **Độ phức tạp**:  
        - ⏱️ Thời gian: `O(n²)`  
        - 💾 Không gian: `O(n)`  
    - **Sorting + Greedy (O(n log n)) – Best Practice**:  
      - Sort intervals theo `start`.  
      - Duyệt tuần tự, so với interval cuối trong `res`:  
        - Nếu overlap → merge lại.  
        - Nếu không → push interval mới vào.  
      - **Độ phức tạp**:  
        - ⏱️ Thời gian: `O(n log n)` (do sort).  
        - 💾 Không gian: `O(1)` (nếu tái sử dụng input).  
  - **Ghi chú**:  
    - Đây là bài cơ bản về **interval problems**.  
    - Pattern thường gặp: **sort + merge**.  
    - Quan trọng cho các bài nâng cao: **Insert Interval, Meeting Rooms, Employee Free Time**.  

---

### Day 10 (2025-09-10)
- **Top75**: Group Anagrams ✅
  - **Ý tưởng**:
    - **Sorting Key (O(n·k log k))**:  
      - Sort từng từ → dùng làm key trong map.  
      - Các từ có key giống nhau thì chung 1 nhóm.  
    - **Frequency Count Key (O(n·k)) – Best Practice**:  
      - Dùng mảng 26 phần tử lưu tần suất ký tự.  
      - Biến thành chuỗi (hashable) để làm key map.  
      - Tốt hơn khi `k` lớn vì không cần sort.  
  - **Độ phức tạp**:
    - Sorting:  
      - ⏱️ Thời gian: `O(n·k log k)`  
      - 💾 Không gian: `O(n·k)`  
    - Counting:  
      - ⏱️ Thời gian: `O(n·k)`  
      - 💾 Không gian: `O(n·k)`  
  - **Ghi chú**:  
    - Đây là bài kinh điển về **hashing + grouping**.  
    - Hai cách thường gặp: **sort** vs **frequency**.  
    - Pattern lặp lại trong nhiều bài: **valid anagram, group shifting strings, alien dictionary**.  

---

### Day 11 (2025-09-11)
- **Top75**: Maximum Product Subarray ✅
  - **Ý tưởng**:
    - **Brute Force (O(n²))**:  
      - Duyệt mọi subarray, tính product và lấy max.  
      - **Độ phức tạp**:  
        - ⏱️ Thời gian: `O(n²)`  
        - 💾 Không gian: `O(1)`  
    - **DP với Max/Min So Far (O(n)) – Best Practice**:  
      - Duy trì `maxSoFar` và `minSoFar` vì số âm có thể lật dấu.  
      - Nếu `num < 0` → swap `maxSoFar` và `minSoFar`.  
      - Update:  
        - `maxSoFar = max(num, num * maxSoFar)`  
        - `minSoFar = min(num, num * minSoFar)`  
      - Kết quả = max của tất cả `maxSoFar`.  
      - **Độ phức tạp**:  
        - ⏱️ Thời gian: `O(n)`  
        - 💾 Không gian: `O(1)`  
  - **Ghi chú**:  
    - Đây là bài đặc trưng của **Dynamic Programming** dạng `tracking max/min`.  
    - Lưu ý case đặc biệt khi gặp số 0 hoặc số âm liên tiếp.  
    - Pattern thường xuất hiện trong các bài: **Maximum Sum Subarray (Kadane)** và các biến thể với tích/tổng.  


---

### Day 12 (2025-09-12)
- **Top75**: Search in Rotated Sorted Array ✅
  - **Ý tưởng**:
    - **Brute Force (O(n))**:  
      - Duyệt từng phần tử, return index nếu tìm thấy.  
      - Không đáp ứng yêu cầu O(log n).  
    - **Binary Search (O(log n)) – Best Practice**:  
      - Mỗi lần chia đôi, ít nhất 1 nửa mảng vẫn sorted.  
      - Kiểm tra target có thuộc nửa sorted hay không → quyết định dịch left/right.  
      - **Độ phức tạp**:  
        - ⏱️ Thời gian: `O(log n)`  
        - 💾 Không gian: `O(1)`  
  - **Ghi chú**:  
    - Đây là bài điển hình áp dụng **Binary Search trên mảng xoay vòng**.  
    - Pattern hữu ích cho nhiều bài dạng:  
      - **Find Minimum in Rotated Sorted Array**  
      - **Search in Rotated Sorted Array II** (có duplicate). 


---

### Day 13 (2025-09-22)
- **Top75**: Reverse Linked List ✅
  - **Ý tưởng**:
    - **Iterative (O(n), O(1))**:
      - Dùng 3 con trỏ: prev, curr, next.
      - Đảo liên kết từng bước.
    - **Recursive (O(n), O(n))**:
      - Đệ quy đến cuối danh sách.
      - Khi quay lui, đảo chiều con trỏ.
  - **Độ phức tạp**:
    - ⏱️ Thời gian: `O(n)`  
    - 💾 Không gian: Iterative `O(1)`, Recursive `O(n)` (stack).  
  - **Ghi chú**:  
    - Đây là bài kinh điển để ôn **Linked List manipulation**.  
    - Interview thường yêu cầu biết cả 2 cách.  

---

### Day 14 (2025-09-23)
- **Top75**: Linked List Cycle ✅
  - **Ý tưởng**:
    - **HashSet (O(n), O(n))**:  
      - Lưu node đã đi qua.  
      - Nếu gặp lại node → có cycle.  
    - **Floyd’s Algorithm (O(n), O(1)) – Best Practice**:  
      - Slow pointer đi 1 bước, fast pointer đi 2 bước.  
      - Nếu gặp nhau → cycle tồn tại.  
      - Nếu fast chạm nil → không có cycle.  
  - **Độ phức tạp**:
    - ⏱️ Thời gian: `O(n)`  
    - 💾 Không gian: HashSet `O(n)`, Floyd `O(1)`  
  - **Ghi chú**:  
    - Bài này cực kỳ quan trọng, là pattern chung để tìm cycle.  
    - Mở rộng:  
      - **Linked List Cycle II** → tìm node bắt đầu cycle.  

---


### Day 15 (2025-09-24)
- **Top75**: Container With Most Water ✅
  - **Ý tưởng**:
    - **Brute Force (O(n²))**:  
      - Thử tất cả cặp `(i, j)`, tính area = min(height[i], height[j]) * (j - i).  
      - Quá chậm với n=1e5.  
    - **Two Pointers (O(n)) – Best Practice**:  
      - Bắt đầu từ 2 đầu mảng (left=0, right=n-1).  
      - Tính area, cập nhật max.  
      - Di chuyển con trỏ có height nhỏ hơn vào trong.  
  - **Độ phức tạp**:
    - ⏱️ Thời gian: `O(n)`  
    - 💾 Không gian: `O(1)`  
  - **Ghi chú**:  
    - Đây là pattern điển hình: **Two Pointers for optimization**.  
    - Trực giác: dịch con trỏ chiều cao nhỏ để có cơ hội tăng diện tích.  

---


### Day 16 (2025-09-25)
- **Top75**: Find Minimum in Rotated Sorted Array ✅
  - **Ý tưởng**:
    - Mảng đã **sorted + rotated** → luôn có một nửa được sort.
    - So sánh `nums[mid]` với `nums[right]`:
      - Nếu `nums[mid] > nums[right]` → min nằm bên phải.
      - Ngược lại → min nằm bên trái (bao gồm mid).
  - **Độ phức tạp**:
    - ⏱️ Thời gian: `O(log n)`
    - 💾 Không gian: `O(1)`
  - **Ghi chú**:
    - Đây là dạng **Binary Search template** rất phổ biến trong các bài rotated array.
    - Pattern này cũng áp dụng cho LeetCode 33, 81.

---


### Day 17 (2025-09-26)
- **Top75**: Longest Repeating Character Replacement ✅
  - **Ý tưởng**:
    - Dùng **sliding window** + đếm tần suất ký tự.
    - Luôn giữ `maxCount` = số lần xuất hiện nhiều nhất trong cửa sổ.
    - Điều kiện hợp lệ: `(window size - maxCount) <= k`.
    - Nếu không hợp lệ → thu nhỏ cửa sổ từ bên trái.
  - **Độ phức tạp**:
    - ⏱️ Thời gian: `O(n)`
    - 💾 Không gian: `O(1)` (26 ký tự uppercase).
  - **Ghi chú**:
    - Pattern sliding window tương tự bài **Longest Substring Without Repeating Characters**.
    - Chú ý **maxCount không giảm** khi thu hẹp cửa sổ → nhưng vẫn đúng, vì kết quả chỉ cần tìm độ dài lớn nhất.


--- 

### Day 18 (2025-09-29)
- **Top75**: Longest Substring Without Repeating Characters ✅
  - **Ý tưởng**:
    - Dùng **sliding window + map** lưu vị trí xuất hiện cuối cùng.
    - Khi gặp ký tự trùng → dịch `left` sang phải.
    - Cập nhật `maxLen`.
  - **Độ phức tạp**:
    - ⏱️ O(n) thời gian.
    - 💾 O(k) không gian (k = kích thước bảng mã, ASCII ≤ 128).
  - **Ghi chú**:
    - Đây là bài **khởi đầu** cho rất nhiều sliding window pattern.
    - Quan trọng: `left` chỉ tăng, không bao giờ giảm.

---

### Day 19 (2025-09-30)
- **Top75**: Number of Islands ✅
  - **Ý tưởng**:
    - Duyệt từng ô trong grid.
    - Khi gặp `'1'`, coi như bắt đầu một đảo mới → DFS/BFS lan ra đánh dấu visited.
    - Đếm số lần khởi động DFS → chính là số đảo.
  - **Độ phức tạp**:
    - ⏱️ O(m × n) thời gian.
    - 💾 O(m × n) không gian (stack/queue).
  - **Ghi chú**:
    - Bài này là nền tảng cho mọi bài **connected components** trong graph.
    - Quan trọng để hiểu pattern DFS/BFS trên ma trận.

---

### Day 20 (2025-10-01)
- **Top75**: Remove Nth Node From End of List ✅
  - **Ý tưởng**:
    - Dùng **dummy node** để tránh case xoá head.
    - Dùng **two pointers**: fast đi trước n+1 bước → sau đó cùng chạy với slow.
    - Khi fast hết → slow ngay trước node cần xoá.
  - **Độ phức tạp**:
    - ⏱️ O(L) với L = độ dài list.
    - 💾 O(1) không gian.
  - **Ghi chú**:
    - Đây là bài kinh điển để luyện kỹ thuật two pointers trong linked list.
    - Dummy node giúp code clean, tránh if else khi xoá head.

---

### Day 21 (2025-10-02)
- **Top75**: Palindromic Substrings ✅
  - **Ý tưởng**:
    - Dùng kỹ thuật **Expand Around Center**.
    - Có `2n - 1` tâm: n cho palin lẻ, n-1 cho palin chẵn.
    - Mỗi lần mở rộng kiểm tra ký tự trái phải có bằng nhau không.
  - **Độ phức tạp**:
    - ⏱️ O(n^2), vì mỗi tâm mở rộng tối đa n.
    - 💾 O(1).
  - **Ghi chú**:
    - Đây là một dạng “trick” để đếm palindrome mà không cần DP table.
    - Nếu muốn, có thể làm bằng **DP[i][j]**: true nếu substring i..j palindrome → cũng O(n^2) nhưng tốn O(n^2) space.

---

### Day 22 (2025-10-03)
- **Top75**: Pacific Atlantic Water Flow ✅
  - Ý tưởng: thay vì đi từ mỗi ô ra biển → đi từ biển vào trong.
  - Dùng DFS/BFS đánh dấu 2 tập (Pacific, Atlantic).
  - Lấy giao 2 tập → kết quả.
- Complexity:
  - ⏱️ O(m*n)
  - 💾 O(m*n)
- Ghi chú:
  - Cách đảo ngược tư duy (reverse flow) là chìa khóa.

---

### Day 23 – Minimum Window Substring ✅ (2025-10-06)
-  **Top75**: Sliding Window  
- **Ý tưởng chính**:
  - Duyệt chuỗi `s` bằng hai con trỏ `left` – `right`.
  - Giữ một **bản đồ tần suất (map)** cho các ký tự cần trong `t`.
  - Khi cửa sổ hiện tại chứa đủ ký tự của `t`, thử thu nhỏ để tìm **cửa sổ nhỏ nhất**.
- **Độ phức tạp**:
  - ⏱ O(m + n)
  - 💾 O(1) (do chỉ dùng map ký tự ASCII)
- **Ghi chú**:
  - Đây là bài kinh điển trong nhóm “Sliding Window Hard”.
  - Hiểu rõ cơ chế “have / need” rất quan trọng.
  - Là nền tảng cho các bài như:
    - `Permutation in String`
    - `Minimum Window Subsequence`

---

### Day 24 (2025-10-07)
- **Top75**: Invert Binary Tree ✅  
  - Bài kinh điển về **DFS + recursion**.  
  - Ý tưởng chính: đảo vị trí left ↔ right ở mỗi node.  
  - Giúp nắm chắc cách duyệt và thao tác trên cấu trúc cây.
- Complexity:
  - ⏱️ O(n)
  - 💾 O(h)
- Ghi chú:
  - Dễ nhưng quan trọng, nền tảng cho các bài như Symmetric Tree, Mirror Tree.

---

### Day 25 (2025-10-08)
* **Top75**: Validate Binary Search Tree ✅
  * Dùng **DFS + giới hạn giá trị (min, max)** để kiểm tra tính hợp lệ.
  * Mỗi node phải nằm trong phạm vi `(min, max)` tùy theo vị trí trong cây.
  * Gợi nhớ bài toán kiểm tra cấu trúc ràng buộc trong cây.
* Complexity:
  * ⏱️ O(n)
  * 💾 O(h)
* Ghi chú:
  * Bài này rất hay để hiểu rõ bản chất của cây tìm kiếm nhị phân.

---

### Day 26 (2025-10-09)
* **Top75**: Non-overlapping Intervals ✅
  * Kỹ thuật: **Greedy chọn khoảng kết thúc sớm nhất**.
  * Ý tưởng giống bài “Activity Selection Problem” trong thuật toán cổ điển.
  * Sắp xếp theo `end time` để tối ưu số lượng giữ lại.
* Complexity:
  * ⏱️ O(n log n)
  * 💾 O(1)
* Ghi chú:
  * Đây là một bài rất hay để luyện tư duy greedy tối ưu hoá lựa chọn.

---

### Day 27 (2025-10-10)
* **Top75**: Construct Binary Tree from Preorder and Inorder Traversal ✅
  * Tư duy chính: dùng **preorder để chọn root**, và **inorder để chia đôi cây**.
  * Duy trì con trỏ `preIdx` trong preorder, tránh tạo mảng con để tối ưu.
* Complexity:
  * ⏱️ O(n)
  * 💾 O(n)
* Ghi chú:
  * Đây là một trong những bài cơ bản và quan trọng nhất khi làm việc với cấu trúc cây nhị phân.

---

### Day 28 (2025-10-13)
* **Top75**: Top K Frequent Elements ✅
  * Cách tiếp cận chính: **HashMap + MinHeap (Priority Queue)**
  * Tối ưu hơn so với sort toàn bộ mảng.
* Complexity:
  * ⏱️ O(n log k)
  * 💾 O(n)
* Ghi chú:
  * Có thể cải tiến lên O(n) bằng **bucket sort** khi không yêu cầu thứ tự cụ thể.
  * Đây là bài kinh điển giúp luyện kỹ năng thao tác với heap trong Go.

---

* ### Day 29 (2025-10-14)
* **Top75**: Clone Graph ✅
  * Dùng DFS (hoặc BFS) để sao chép toàn bộ cấu trúc đồ thị.
  * Cần chú ý xử lý **chu trình (cycle)** bằng `visited` map.
* Complexity:
  * ⏱️ O(N + E)
  * 💾 O(N)
* Ghi chú:
  * Đây là bài tiêu biểu kiểm tra tư duy “deep copy” trong cấu trúc dữ liệu có liên kết vòng.
  * Nên thử cả hai cách DFS và BFS để nắm rõ sự khác biệt về triển khai.

---

* ### Day 30 (2025-10-15)
* **Top75**: Course Schedule ✅
  * Mô hình đồ thị có hướng, phát hiện chu trình bằng DFS.
  * Nếu có chu trình → không thể hoàn thành tất cả khóa học.
* Complexity:
  * ⏱️ O(V + E)
  * 💾 O(V + E)
* Ghi chú:
  * Bài rất cơ bản để hiểu **Topological Sort**.
  * Nên luyện cả 2 cách: DFS và BFS (Kahn’s Algorithm) để so sánh hiệu quả.