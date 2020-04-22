1. `start`  
&nbsp;&nbsp;**if** SUCCESS **then**:  
&nbsp;&nbsp;&nbsp;&nbsp;`stop`  
&nbsp;&nbsp;**else**  
&nbsp;&nbsp;&nbsp;&nbsp;**if** VM does not exist **then**  
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;`delete` or `start`  
&nbsp;&nbsp;&nbsp;&nbsp;**else**  
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;**if** VM off   
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;**if** 'everything OK' **then**    // which conditions imply 'OK'  
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;`start` or `delete`  
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;**else**  
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;`delete`  
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;**else** // VM on  
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;`stop`  
  				  
2. `stop`  
&nbsp;&nbsp;**if** SUCCESS **then**  
&nbsp;&nbsp;&nbsp;&nbsp;`start` or `delete`  
&nbsp;&nbsp;**else**  
&nbsp;&nbsp;&nbsp;&nbsp;**if** VM off **then**  
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;`delete`   // no start because of failed stop  
&nbsp;&nbsp;&nbsp;&nbsp;**else**  
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;`stop -f`   // instead of stop, otherwise cyclic  

3. `delete`  
&nbsp;&nbsp;**if** SUCCESS **then**  
&nbsp;&nbsp;&nbsp;&nbsp;`start`  
&nbsp;&nbsp;**else**  
&nbsp;&nbsp;&nbsp;&nbsp;**if** VM exists **then**  
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;`start` or `delete -f`
&nbsp;&nbsp;&nbsp;&nbsp;**else**  
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;`delete -f`   // instead of delete, otherwise cyclic  

