 function retroceso(){
     const table = document.getElementById("mytable");
     const fila = table.rows[1];
     const celda = fila.cells[1];
     const num = celda.textContent;
     if(parseInt(num)>1){
        window.location.href = "/superhero?input_busqueda=" + (parseInt(num)-1).toString()
     }
 }

 function avanzar(){
     const table = document.getElementById("mytable");
     const fila = table.rows[1];
     const celda = fila.cells[1];
     const num = celda.textContent;
     if(parseInt(num)<732){
        window.location.href = "/superhero?input_busqueda=" + (parseInt(num)+1).toString()
     }
 }