// Global site tag (gtag.js) - Google Analytics
(function () {
  var id = 'UA-97015161-1';
  var gcse = document.createElement('script');
  gcse.type = 'text/javascript';
  gcse.async = true;
  gcse.src = 'https://www.googletagmanager.com/gtag/js?id=' + id;
  var s = document.getElementsByTagName('script')[0];
  s.parentNode.insertBefore(gcse, s);

  window.dataLayer = window.dataLayer || [];
  function gtag(){dataLayer.push(arguments);}
  gtag('js', new Date());

  gtag('config', id);
})();