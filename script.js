document.addEventListener('DOMContentLoaded', function() {
	const openmenu = document.getElementById('openmenu')
	const closemenu = document.getElementById('closemenu')
	const main = document.getElementById('main')
	const body = document.body
	const side = document.getElementById('sidebar')

	openmenu.addEventListener('click', function() {
		body.classList.add('bg-gray-800', 'bg-opacity-70');
		side.classList.toggle('translate-x-full');
	})

	closemenu.addEventListener('click', function() {
		body.classList.remove('bg-gray-800', 'bg-opacity-70');
		side.classList.toggle('translate-x-full');
	})
}
)
