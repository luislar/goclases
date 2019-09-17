var url = "ws://" + window.location.host + "/ws";
var ws = new WebSocket(url);
var datos = document.getElementById("datos");
indice=1;

ws.onmessage = function (msg) {
   orden = msg.data.split("#");
   $(".datos").append("<h6 class='ani c"+indice.toString()+"'>* "+orden[1]+"</h6>");
  
   $(".c"+indice.toString()).animateCss('zoomIn');
   indice++;
}

//Se crea la funcion "animateCss" para animar un texto cuando este aparece.
$.fn.extend({
    animateCss: function (animationName, callback) {
      var animationEnd = 'webkitAnimationEnd mozAnimationEnd MSAnimationEnd oanimationend animationend';
      this.addClass('animated ' + animationName).one(animationEnd, function() {
          $(this).removeClass('animated ' + animationName);
          if (callback) {callback();}
      });
    return this;
  }});