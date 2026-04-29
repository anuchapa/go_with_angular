import { Component, input, output } from '@angular/core';

@Component({
  selector: 'app-products-list',
  imports: [],
  templateUrl: './products-list.html',
  styleUrl: './products-list.css',
})
export class ProductsList {
  public readonly products = input<any[]>()
  public readonly deleteEvent = output<number>()
  public readonly openQrEvent = output<string>()
  public readonly openDelEvent = output<number>()


  // public deleteProduct(index: number) {
  //   this.deleteEvent.emit(index)
  // }

  public openQR(productCode: string) {
    this.openQrEvent.emit(productCode)
  }

  public openDelConfirm(index: number) {
    this.openDelEvent.emit(index)
  }
}
