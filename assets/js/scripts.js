var i = 0;

$(init);
function init() {
    $('#addOne').bind('click', add_one);
    $('#showCount').bind('click', show_count);
}

function add_one() {
    $('#counterButtons')
        .append('<div id="counter" class="animated textOut">+1</div>');
    $('#counter').attr("id", 'counter' + i);
    $('#counter'+i)
        .fadeIn(1000)
        .fadeOut(2000, function() {
            $(this).remove()
        });

    $.ajax({
        url: '/add',
        type: 'POST',
        data: [],
        dataType: 'json'
    });

    i++;
    if (i > 500) {
        i = 0;
    }
}

function show_count() {
    $("#result").remove();
    $.ajax({
        url: '/show',
        type: 'GET',
        success: function(count){
            count = count.toString();
            $("#currentCount")
				.append('<div id="result">Текущее значение счетчика = ' + count + '</div>');

            $("#result").fadeIn(2000).fadeOut(4000, function(){$(this).remove()});
        }
    });
}