
let data = [{
    name,
    desc,

}]

function handleForm(event){
    // event.preventDefault
    var name = document.getElementById('name').value
    var desc = document.getElementById('desc').value

    // console.log( name, desc)

    data.push({name:name,desc,desc})

    for(index=0;index<data;data++){
        console.log(data);
    }    


}