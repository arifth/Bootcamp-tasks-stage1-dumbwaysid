let data =[]

function handleForm(event){
    let datablog= []

    // use obj desctructuring here 
    // https://www.freecodecamp.org/news/array-and-object-destructuring-in-javascript/
    
    let data = { name: "", start:"", end:"", desc:"", img:"" }

    // event.preventDefault
    let name = document.getElementById('name').value
    let start = document.getElementById('startdate').value
    let end = document.getElementById('enddate').value
    let desc = document.getElementById('desc').value
    let img = document.getElementById('input-img').value

    console.log( name, start, end, desc, img)

    let blogData = 

    datablog.push({name , start, end, desc, img})

    console.log(data);

    // for(index=0;index<data;data++){
    //     console.log(data);
    // }    


}