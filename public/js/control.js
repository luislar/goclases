
var url = "ws://" + window.location.host + "/ws";
var ws = new WebSocket(url);
var inicio = 1;

function Siguiente(){
    inicio++;
    //Enviar("Siguiente");
    $(".titulos").val(inicio);
}

function Anterior(){
    inicio--;
    $(".titulos").val(inicio);
}

function Enviar(mensaje){
    ws.send(inicio.toString()+"#"+mensaje);
}

const TraerTitulos = async () => {
  const response = await fetch('http://'+ window.location.host +'/titulos');
  const js = await response.json();
  for (index = 0; index < js.length; index++) {
  	opciones = "<option value='"+js[index].Id+"' class='th2'>"+js[index].Titulo+"</option>";
    $(".titulos").append(opciones);
   }
}



TraerTitulos();
