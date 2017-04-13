# HTML Input hidden defaultValue vs. value with jQuery val() vs. attr()

## Problem

承續上一篇 [HTML Input defaultValue vs. value, jQuery val() vs. attr()](https://dinos80152.wordpress.com/2015/10/07/html-input-defaultvalue-vs-value-jquery-val-vs-attr/)

這篇探討為什麼 type 使用 hidden，html 就成功的被改變了呢？

## Example

範例如下

```html
<form>
<input type="hidden" name="test" value="abc">
<input type="hidden" name="test2" value="abc">
</form>
```

Javascript

```js
document.getElementsByName('test')[0].value = 123;
document.getElementsByName('test')[0].setAttribute('value', 123);
console.log(document.getElementsByTagName('form')[0].innerHTML);
```

jQuery

```js
$('input[name="test"]').val(123);
$('input[name="test"]').attr('value', 123);
console.log($('form').html());
```

Output

```html
<input type="text" name="test" value="123">
<input type="text" name="test2" value="123">
```

原因很簡單，一樣參考[Input 的屬性](http://www.w3.org/TR/DOM-Level-2-HTML/html.html#ID-6043025)。

我們可以看到 defaultValue 是只有 **text**, **file**, **password** 才有的屬性，所以 hidden 這個 type 並沒有 defaultValue，只有 value，因為在制定時，hidden 的值是不會被 user 改變的。

因為沒有 defaultValue，所以 input type hidden 上原本的 value attribute 就是對到 value property，所以當你使用 js 改變 value 時，它的 html 也會一起被改變了。

但這邊最奇怪的地方，就算它沒有 defaultValue 這個屬性，但依舊可以使用

```js
document.getElementsByName('test')[0].setAttribute('value', 123);
document.getElementsByName('test')[0].defaultValue = 123;
```

這兩個 method 則會去改 input type=hidden 元素的 value 了。

這該說方便還是讓人混淆呢..?

或許之後有機會看到更多討論再來做更多的釐清吧！

### Reference

* [Javascript and defaultValue of hidden input elements@stackoverflow](http://stackoverflow.com/questions/5319678/javascript-and-defaultvalue-of-hidden-input-elements)
* [Firefox will not reset hidden fields@mozillazine](http://forums.mozillazine.org/viewtopic.php?f=25&t=1787115)
* [Document Object Model HTML-Interface HTMLInputElement](http://www.w3.org/TR/DOM-Level-2-HTML/html.html#ID-6043025)