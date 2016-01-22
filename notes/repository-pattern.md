# Repository
源自 c#
repository: 知識庫

## 好處
* 隔離資料層
  * 方便單元測試
  * 利於將來可抽換 (ex: 從 Relational Database 換到 NoSQL or API)

## Example

Controllers/NewsController.php
```
class NewsController extends Controller
{
    public function index() {
        return News::listAllAble();
    }
}
```

Models/News.php
```
class News extends Model
{
    public function listAllAble() {
        return $this->where('status', 1)->get();
    }
}
```

Repositories/NewsRepository.php
```
use App\Models\News;

class NewsRepository
{
    public function listAllAble() {
        return News::where('status', 1)->get();
    }
}
```



## Reference
[The Repository Pattern](https://msdn.microsoft.com/en-us/library/ff649690.aspx)
[Repository，我可能不會用你](http://huan-lin.blogspot.com/2012/11/repository-yagni.html)
