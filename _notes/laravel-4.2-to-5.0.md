# Laravel Update 4.2 to 5.0

## Basic Flow

```sequence
/public/index.php->Route:
Route->Controller:
Controller->Model:
Model-->Controller:data
Controller->View:
```

## Env

**Laravel 5 doesn't have environment detection **, use .env.example

|     | 4.2                    | 5    |
| --- | ---------------------- | ---- |
| env | app/config/[envfolder] | .env |

## Namespace

* Naming Your Application (optional)

    ```bash
    php artisan app:name Mstar
    ```

## Route

|               | 4.2           | 5                                      |
| ------------- | ------------- | -------------------------------------- |
| middleware    | filter        | middleware                             |
| where(global) | app/route.php | app/Providers/RouteServiceProvider.php |

### Middleware

**CSRF Protection is default middleware in Laravel 5**

* 4.2

    ```php
    // app/filter.php
    Route::filter('csrf', function()
    {
        if (Session::token() != Input::get('_token'))
        {
            throw new IlluminateSessionTokenMismatchException;
        }
    });
    ```

* 5

    ```php
    // app/Http/Middleware/VerifyCsrfToken.php
    class VerifyCsrfToken extends BaseVerifier {

        public function handle($request, Closure $next)
        {
            return parent::handle($request, $next);
        }
    }
    ```

#### Register

```php
// app/Http/Kernel.php

// The application's global HTTP middleware stack.
protected $middleware = [
    //...
    'AppHttpMiddlewareVerifyCsrfToken',
];

// The application's route middleware.
protected $routeMiddleware = [
    //...
    'csrf' => 'AppHttpMiddlewareVerifyCsrfToken',
];
```

#### Using

* 4.2

    ```php
    Route::get('/', array('before' => 'csrf|auth', 'uses' => 'index@HomeController'));
    ```

* 5

    ```php
    Route::get('/', ['middleware' => ['csrf', 'auth'], 'uses' => 'index@HomeController']);
    ```

### Where

#### Global

* 4.2

    ```php
    // app/route.php
    Route::pattern('id', '[0-9]+');
    ```

* 5

    ```php
    // app/Providers/RouteServiceProvider.php
    public function boot(Router $router)
    {
        $router->pattern('id', '[0-9]+');
        parent::boot($router);
    }
    ```

## Controller

### Parent Class & Namespace

* 4.2

    ```php
    // app/controllers/HomeController.php

    class HomeController extends BaseController {}
    ```

* 5

    ```php
    // app/Http/Controllers/HomeController.php

    namespace AppHttpControllers;

    class HomeController extends Controller {}
    ```

## Model

### Parent Class & Namespace

* 4.2

    ```php
    // app/models/User.php
    class User extends Eloquent {
    }
    ```

* 5

    ```php
    // app/User.php
    use IlluminateDatabaseEloquentModel;

    class User extends Model {
    }
    ```

### Others

* Laravel 5: if you wanna create instance, you have to set

    ```php
    protected $guarded = [];

    //default
    protected $guarded = ['*'];
    ```

* Relation has to setting namespace

    ```php
    public function users()
    {
        return $this->hasMany('AppModelsUser');
    }
    ```

### Eager Loading

* 4.2

    ```php
    $report->report_assigns
    ```

* 5

    ```php
    $report->reportAssigns
    ```

## View

### Echo

|        | 4.2             | 5               |
| ------ | --------------- | --------------- |
| escape | `{{{ $data }}}` | `{{ $data }}`   |
| raw    | `{{ $data }}`   | `{!! $data !!}` |

### HTML Helper

* 4.2
  * HTML::
  * FORM::
* 5 (It's removed, solution)
  * [Adding html package](http://laravelcollective.com/docs/5.0/html)
  * Using original html tag

    ```html
    <script src="{{ asset("js/jquery.js") }}"></script>
    ```

    ```html
    <link rel="stylesheet" href="{{ asset("css/style.css") }}">
    ```

    ```html
    <form>
        <input type="hidden" name="_token" value="{{ csrf_token() }}">
        <input type="hidden" name="_method" value="PUT">
    </form>
    ```

## Pagination

### Methods

| 4.2              | 5             |
| ---------------- | ------------- |
| **getTotal()**   | **total()**   |
| **links()**      | **render()**  |
| getCurrentPage() | currentPage() |
| getLastPage()    | lastPage()    |
| getPerPage()     | perPage()     |
| getFrom()        | firstItem()   |
| getTo()          | lastItem()    |

### Customize

* 4.2, modify config or dynamic setting

    ```php
    // app/config/view.php
    'pagination' => 'pagination::slider-3',
    ```

    ```php
    // views/news.php
    $news->links('view.name');
    ```

* 5
    1. Implement Presenter

        ```php
        // Illuminate/Pagination/BootstrapThreePresenter.php

        use IlluminateContractsPaginationPresenter as PresenterContract;

        class BootstrapThreePresenter implements PresenterContract
        {
        }
        ```

    1. Give render method presenter object

        ```php
        $news->render($bootstrapThreePresenter);
        ```

## Error Handler

| 4.2                  | 5               |
| -------------------- | --------------- |
| app/start/global.php | app/Exceptions/ |

## Package

### Auth

* 4.2, it doesn't maintain anymore.

    ```json
    "artdarek/oauth-4-laravel"
    ```

* 5 (Official Package: Socialite), **Socialite needs your Google app to have Google+ API enabled.**

    ```json
    "laravel/socialite": "~2.0"
    ```

### Predis

* Adding predis to proeject composer.json require block, **In laravel 5,  predis move from require to require-dev.**

    ```json
    "predis/predis": "~1.0"
    ```

## Unit Test

| diff     | 4.2                | 5           |
| -------- | ------------------ | ----------- |
| test env | app/config/testing | phpunit.xml |

## Core Concept

### Service Providers

#### Folder

```bash
app/Providers
```

#### Register

```php
// app/config/
'providers' => [
    //...
]
```

### Contracts

### Repositories

#### Binding Interfaces To Implementations

```php
// app/Providers/AppServiceProvider.php
public function register()
{
    $this->app->bind(
        'App\Contracts\News',
        'App\Repsitories\EloquentNewsRepository'
    );
}
```

**Remember to Clear Compiled Cache**

```bash
php artisan clear-compiled
```

## Facades