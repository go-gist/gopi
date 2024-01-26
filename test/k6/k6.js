import http from 'k6/http';
import { check, sleep } from 'k6';

export const options = {
  vus: 1, // Number of virtual users
  duration: '30s', // Duration of the test
  thresholds: {
    // Define thresholds for response time (in milliseconds)
    http_req_duration: ['p(95)<200', 'p(99)<300'],

    // Define thresholds for throughput (requests per second)
    http_reqs: ['rate>100'],

    // Define thresholds for error rate
    http_req_failed: ['rate<0.1'],
  },
  ext: {
    loadimpact: {
      projectID: 123456, // Replace with your Load Impact project ID
      // Add any Load Impact specific configurations
    },
  },
  // Output the response bodies to a file
  noConnectionReuse: true,
  discardResponseBodies: true,
};

export default function () {
  const id = 1;
  const url = `http://192.168.1.17:8080/foo?id=${id}`;
  const payload = JSON.stringify({
    key1: 'value1',
    key2: 'value2',
  });
  const headers = {
    'Content-Type': 'application/json',
  };

  const response = http.post(url, payload, { headers });

  // Check if the request was successful (HTTP status 200)
  check(response, {
    'is successful': (r) => r.status === 200,
  });

  // Log the response body to the console
  console.log(`Response body: ${response.body}`);

  // Sleep for a short duration between requests if needed
  sleep(1);
}
