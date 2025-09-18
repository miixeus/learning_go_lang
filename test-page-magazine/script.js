$(document).ready(function () {
    let currentPage = 0;
    const pages = $('.page');
    const totalPages = pages.length;

    // Função para atualizar a posição da "revista" (deslizar entre as páginas)
    function updateBookPosition() {
        const offset = -currentPage * 100; // Desloca 100% da largura por página
        $('.book').css('transform', 'translateX(' + offset + 'vw)');
    }

    // Ação ao clicar no botão "Próxima"
    $('#next').click(function () {
        if (currentPage < totalPages - 1) {
            currentPage++;
            updateBookPosition();
        }
    });

    // Ação ao clicar no botão "Anterior"
    $('#prev').click(function () {
        if (currentPage > 0) {
            currentPage--;
            updateBookPosition();
        }
    });

    // Ação ao clicar nos links da barra de navegação
    $('.page-link').click(function () {
        currentPage = $(this).data('page');
        updateBookPosition();
    });

    // Função para navegação com as setas do teclado
    $(document).keydown(function (e) {
        if (e.key === "ArrowRight" && currentPage < totalPages - 1) {
            currentPage++;
            updateBookPosition();
        } else if (e.key === "ArrowLeft" && currentPage > 0) {
            currentPage--;
            updateBookPosition();
        }
    });

    // Função para permitir navegação por toque
    let touchStartX = 0;

    $('.book-container').on('touchstart', function (e) {
        touchStartX = e.originalEvent.touches[0].pageX;
    });

    $('.book-container').on('touchmove', function (e) {
        let touchEndX = e.originalEvent.touches[0].pageX;
        let touchDiff = touchStartX - touchEndX;

        if (touchDiff > 50 && currentPage < totalPages - 1) { // Deslizar para a direita
            currentPage++;
            updateBookPosition();
        } else if (touchDiff < -50 && currentPage > 0) { // Deslizar para a esquerda
            currentPage--;
            updateBookPosition();
        }
    });
});
