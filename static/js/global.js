document.addEventListener('htmx:beforeSwap', function (evt) {
  if (evt.detail.xhr.status === 422) {
    // allow 422 responses to swap as we are using this as a signal that
    // a form was submitted with bad data and want to rerender with the errors
    evt.detail.shouldSwap = true;

    // set isError to false to avoid error logging in console
    evt.detail.isError = false;
  }
});

const hiddenFlexToggler = (element) => {
  element.classList.toggle('hidden');
  element.classList.toggle('flex');
};

const globalModalToggle = (event, element) => {
  let modal = document.getElementById(element.dataset.modalId);
  hiddenFlexToggler(modal);

  event.stopPropagation();
};
