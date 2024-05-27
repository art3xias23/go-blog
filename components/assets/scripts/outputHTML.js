document.addEventListener('DOMContentLoaded', function() {
    var textarea = document.getElementById('myContent');
    var outputDiv = document.getElementById('myOutput');

    if(textarea && outputDiv) {
        textarea.addEventListener('input', function() {
            outputDiv.innerHTML = textarea.value;
        });
    }
});
