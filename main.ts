class Order {
  price : number;

  basePrice() {
    return this.price;
  }
}

function badHasDiscount(order: Order): boolean {
  let basePrice: number = order.basePrice();
  return basePrice > 1000;
}

function hasDiscount(order: Order): boolean {
  return order.basePrice() > 1000;
}

function discountAmount(order: Order): number {
  if (order.price > )
}





class BadAccount {
  sensitiveData: string;
}

class Account {
  private _sensitiveData: string;
  seniority : number;
  isPartTime : number;
  isPrivileged: boolean;
  get sensitiveData() {
    return this._sensitiveData;
  }
  setSensitiveData(arg: string): void {
    this._sensitiveData = arg;
  }

  isEligibleForDiscount() : boolean{
    return true;
  }

  discountAmount(): number {
    if (!this.isEligibleForDiscount) {
      return 0;
    }
    //...
  }
}

let bb : Account = new Account();

bb.setSensitiveData()
bb.discountAmount()
bb.sensitiveData

function badDisabilityAmount(years: number): number {
  if (seniority < 2) {
    return 0;
  }
  if (monthsDisabled > 12) {
    return 0;
  }
  if (isPartTime) {
    return 0;
  }
  // Compute the disability amount.
  // ...
}

function disabilityAmount(years: number): number {
  if (isNotEligibleForDisability()) {
    return 0;
  }
  // Compute the disability amount.
  // ...
}
