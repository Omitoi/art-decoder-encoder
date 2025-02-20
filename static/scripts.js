function adjustFontSize(elementId1, elementId2, increment) {
    const element1 = document.getElementById(elementId1);
    const element2 = document.getElementById(elementId2);
    if (element1) {
        const computedStyle1 = window.getComputedStyle(element1, null);
        const style1 = computedStyle1.getPropertyValue('font-size');
        const currentSize1 = parseFloat(style1);
        element1.style.fontSize = (currentSize1 + increment) + 'px';
    }
    if (element2) {
        const computedStyle2 = window.getComputedStyle(element2, null);
        const style2 = computedStyle2.getPropertyValue('font-size');
        const currentSize2 = parseFloat(style2);
        element2.style.fontSize = (currentSize2 + increment) + 'px';
    }
}

function autoResizeTextarea(textareaId) {
    const textarea = document.getElementById(textareaId);
    if (textarea) {
        textarea.style.height = 'auto';
        textarea.style.height = textarea.scrollHeight + 'px';
    }
}


document.addEventListener('DOMContentLoaded', function() {
    const textarea = document.getElementById('art_input');
    textarea.addEventListener('input', function() {
        autoResizeTextarea('art_input');
    });
    autoResizeTextarea('art_input'); // Initial resize
});