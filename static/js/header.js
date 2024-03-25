document.addEventListener('DOMContentLoaded', () => {
  const toggleDarkModeCheck = document.querySelector('#darkModeToggle');

  toggleDarkModeCheck.addEventListener('change', (e) => {
    component = document.body;
    component.classList.toggle('dark');
    component.classList.toggle('text-bodydark');
    component.classList.toggle('bg-boxdark-2');
    e.stopPropagation();
  });

  const userMenuLink = document.querySelector('#userMenuLink');
  const userMenuDropdown = document.querySelector('#userMenuDropdown');

  userMenuLink.addEventListener('click', (e) => {
    userMenuDropdown.classList.toggle('hidden');
    e.stopPropagation();
  });

  document.addEventListener('click', () => {
    userMenuDropdown.classList.add('hidden');
  });

  // Sidebar state

  var sideBarShowing = false;
  const toggleSideBar = document.querySelector('#toggleSideBar');
  const navbar = document.querySelector('#navbar');
  const toggleSideBarHamburger = document.querySelector('#toggleSideBarHamburger');
  const sideBarHamburgerSpan = document.querySelector('#sideBarHamburger');

  toggle = () => {
    sideBarShowing = !sideBarShowing;

    if (sideBarShowing) {
      navbar.classList.add('-translate-x-full');
      navbar.classList.remove('translate-x-0');
    } else {
      navbar.classList.add('translate-x-0');
      navbar.classList.remove('-translate-x-full');
    }

    sideBarHamburgerSpan.classList.toggle('show-sidebar');
  };

  toggleSideBarHamburger.addEventListener('click', () => {
    toggle();
  });

  toggleSideBar.addEventListener('click', () => {
    toggle();
  });
});
