facuud2-linux@Facuud2-PC:/mnt/f/Proyectos programacion/Folkode-Projects/mediapp/backend$ k6 run login_test.js

         /\      Grafana   /‾‾/
    /\  /  \     |\  __   /  /
   /  \/    \    | |/ /  /   ‾‾\
  /          \   |   (  |  (‾)  |
 / __________ \  |_|\_\  \_____/

     execution: local
        script: login_test.js
        output: -

     scenarios: (100.00%) 1 scenario, 50 max VUs, 2m10s max duration (incl. graceful stop):
              * default: Up to 50 looping VUs for 1m40s over 3 stages (gracefulRampDown: 30s, gracefulStop: 30s)



  █ THRESHOLDS

    http_req_duration
    ✗ 'p(95)<500' p(95)=841.55ms

    http_req_failed
    ✓ 'rate<0.01' rate=0.00%


  █ TOTAL RESULTS

    checks_total.......: 5424    53.685279/s
    checks_succeeded...: 100.00% 5424 out of 5424
    checks_failed......: 0.00%   0 out of 5424

    ✓ login: status is 200
    ✓ login: response has refresh token

    HTTP
    http_req_duration..............: avg=489.12ms min=412.11ms med=443.52ms max=4.13s p(90)=511.45ms p(95)=841.55ms
      { expected_response:true }...: avg=489.12ms min=412.11ms med=443.52ms max=4.13s p(90)=511.45ms p(95)=841.55ms
    http_req_failed................: 0.00%  0 out of 2712
    http_reqs......................: 2712   26.84264/s

    EXECUTION
    iteration_duration.............: avg=1.48s    min=1.41s    med=1.44s    max=5.13s p(90)=1.51s    p(95)=1.84s
    iterations.....................: 2712   26.84264/s
    vus............................: 1      min=1         max=50
    vus_max........................: 50     min=50        max=50
    iterations.....................: 2712   26.84264/s
    vus............................: 1      min=1         max=50
    vus_max........................: 50     min=50        max=50

    NETWORK
    data_received..................: 2.6 MB 25 kB/s
    data_sent......................: 461 kB 4.6 kB/s




running (1m41.0s), 00/50 VUs, 2712 complete and 0 interrupted iterations
default ✓ [======================================] 00/50 VUs  1m40s
ERRO[0101] thresholds on metrics 'http_req_duration' have been crossed