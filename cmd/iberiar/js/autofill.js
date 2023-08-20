function getParameterValue(parameterName) {
    const urlParams = new URLSearchParams(window.location.search);
    return urlParams.get(parameterName);
}

function handleAutofill() {
    const autofillElements = document.querySelectorAll('[js-autofill="true"]');
    
    autofillElements.forEach(function(element) {
        const parameterName = element.getAttribute('name');
        if (parameterName) {
            let parameterValue = getParameterValue(parameterName);
            
            // If parameterValue is not found in GET parameters, check POST data
            if (!parameterValue && document.forms.length > 0) {
                const formData = new FormData(document.forms[0]);
                parameterValue = formData.get(parameterName);
            }
            
            if (parameterValue !== null) {
                if (element.tagName === 'INPUT' || element.tagName === 'TEXTAREA') {
                    element.value = parameterValue;
                } else {
                    element.textContent = parameterValue;
                }
            }
        }
    });
}

document.addEventListener('htmx:afterRequest', function(evt) {
    const copyFromFields = document.querySelectorAll('[hx-copy-from]');
    copyFromFields.forEach(function(field) {
        const copyFromSelector = field.getAttribute('hx-copy-from');
        const copyFromField = document.querySelector(copyFromSelector);
        
        if (copyFromField) {
            field.value = copyFromField.value;
        }
    });
});


document.addEventListener('DOMContentLoaded', function() {
    handleAutofill();
});
