# design

## root contract

+ proxy

## contract type

+ 1 owner (caller limited)
+ 2 auth (set authrity) can upgrade
+ 3 access (erc control) can upgrade
+ 4 erc (erc data) non-upgrade
+ 5 pool (pledge data) non-upgrade  
+ 6 role (role and group data) non-upgrade 
+ 7 token (token data) non-upgrade
+ 8 pledge (pledge control) can upgrade
+ 9 issue (issue control) can upgrade
+ 10 fs (fs data) non-upgrade


+ 100 control (interact control) can upgrade
+ 101 getter

## base contracts

erc
token
pool
pledge
role
issue
fs

## layer

proxy -> control -> base contracts 

getter -> base contracts 


## deploy

### deploy erc

+ deploy access 
+ deploy erc

### deploy role

+ deploy auth => auth address
+ deploy proxy(auth)  => proxy address as root
+ deploy getter(proxy) => getter address
+ deploy control(proxy, auth) => control address 
+ deploy token(control, auth) => token address
+ deploy pool(control, auth) => pool address
+ deploy pledge(control, auth) => pledge address
+ deploy role(control, auth) => role address

+ set token, pool, pledge, role to control address

### add token

+ addT(erc address)

### create group

+ deploy fs(control, auth, gIndex) => fs address
+ createGroup(level, fsAddr, kpledge, ppledge) => create fs pool(control, auth) => fs pool address
 
### register

+ register account => role index
+ register role => rtype
+ add to group

note: keeper need activate

## golang

+ go-mefs only need proxy.sol and getter.sol; 

## upgrade

+ data sol: better no upgrade
+ control sol: can upgrade 
+ base control sol: can upgrage
+ auth sol: can upgrade, be caution;
