document.addEventListener("DOMContentLoaded", function() {
    const form = document.getElementById('metadataForm');
    const outputArea = document.querySelector('#jsonOutput textarea');
    const downloadLink = document.getElementById('downloadJsonLink'); // Ensure this is defined

    form.addEventListener('submit', function(e) {
        e.preventDefault(); // Prevent the default form submission

        const formData = new FormData(form);
        const contributorsList = formData.get('contributors').split(',').map(item => item.trim()); // Split by comma and trim whitespace

        const metadata = {
            DCMI_Metadata: {
                Title: formData.get('title'),
                Creator: formData.get('creator'),
                Subject: formData.get('subject'),
                Description: formData.get('description'),
                Publisher: formData.get('publisher'),
                Contributors: contributorsList,
                Date: formData.get('date'),
                Type: formData.get('type'),
                Format: formData.get('format'),
                Identifier: formData.get('identifier'),
                Source: formData.get('source'),
                Language: formData.get('language'),
                Relation: formData.get('relation'),
                Coverage: formData.get('coverage'),
                Rights: formData.get('rights'),
                DatasetDivisionDescription: formData.get('datasetDivisionDescription')
                // Assuming CID will be added later in the process, so it's not included here
            }
        };

        // Convert metadata to JSON string and display in the textarea
        const jsonString = JSON.stringify(metadata, null, 2);
        outputArea.value = jsonString;

        // Create a Blob from the JSON string
        const blob = new Blob([jsonString], {type: "application/json"});

        // If a previous blob URL exists, revoke it to free up memory
        if (downloadLink.getAttribute('href')) {
            URL.revokeObjectURL(downloadLink.getAttribute('href'));
        }

        // Generate a new blob URL and update the download link
        const url = URL.createObjectURL(blob);
        downloadLink.href = url;
        downloadLink.style.display = 'inline-block'; // Make the download link visible
    });
});






