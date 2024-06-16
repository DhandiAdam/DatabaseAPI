document.addEventListener('DOMContentLoaded', function () {
    if (document.getElementById('create-item-form')) {
        document.getElementById('create-item-form').addEventListener('submit', function (e) {
            e.preventDefault();

            const name = document.getElementById('name').value;
            const description = document.getElementById('description').value;
            const harga = document.getElementById('harga').value;

            fetch('/items', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json'
                },
                body: JSON.stringify({ name, description, harga })
            })
            .then(response => response.json())
            .then(data => {
                if (data.error) {
                    alert('Error: ' + data.error);
                } else {
                    window.location.href = '/';
                }
            })
            .catch(error => {
                alert('Error: ' + error);
            });
        });
    }

    if (document.getElementById('items-list')) {
        fetch('/items')
            .then(response => response.json())
            .then(data => {
                const itemsList = document.getElementById('items-list');
                data.forEach(item => {
                    const listItem = document.createElement('li');
                    listItem.textContent = `${item.name}: ${item.description}: ${item.harga}`;
                    itemsList.appendChild(listItem);
                });
            })
            .catch(error => {
                alert('Error: ' + error);
            });
    }
});
