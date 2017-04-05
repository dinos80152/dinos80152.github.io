## Folder and File
diff   |4.2 | 5
-------|-----|------
404 error | bootstrap/start.php | app/Exceptions/
env| .env | app/config/[envfolder]

## Model
* laravel 5: if you wanna create instance, you have to set 
```
protected $guarded = [];
```
* relation has to setting namespace
```
public function users()
{
    return $this->hasMany('App\Models\User');
}
```

## IoC

### Providers
app/config/
providers

### Contracts

### Repositories
php artisan clear-compiled

## Services

## Facades

## Unit Test
diff   |4.2 | 5
-------|-----|------
test env | app/config/testing | phpunit.xml

### PHPUNIT
#### phpunit.xml
1. testsuites
2. filter
3. php

#### Test Exception
```
$this->setExpectedException('App\Exceptions\APIException');
```

#### Mock
Mockery

#### Coverage
PHP Code Coverage
https://github.com/sebastianbergmann/php-code-coverage

## Package
### Auth
> "artdarek/oauth-4-laravel" to "laravel/socialite": "~2.0"(official)
> **Socialite needs your Google app to have Google+ API enabled.**

### HTML Helper
remove at Laravel 5
http://laravelcollective.com/docs/5.0/html

### Predis
predis from require to require-dev

## Relations
* 4.2: report->report_assigns
* 5: report->reportAssigns

## Pagination
* getTotal -> total
*  links() -> render()
