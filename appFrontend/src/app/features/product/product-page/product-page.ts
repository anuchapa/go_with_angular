import { Component, ElementRef, inject, signal, ViewChild } from '@angular/core';
import { AddProducts } from '../add-products/add-products';
import { ProductsList } from '../products-list/products-list';
import { ProductApi } from '../../../core/services/api/products/product-api';
import { QRCodeComponent } from 'angularx-qrcode';

@Component({
  selector: 'app-product-page',
  imports: [AddProducts, ProductsList, QRCodeComponent],
  templateUrl: './product-page.html',
  styleUrl: './product-page.css',
})
export class ProductPage {
  @ViewChild('myQrDialog') myQrDialog!: ElementRef<HTMLDialogElement>;
  @ViewChild('myConfirmDialog') myConfirmDialog!: ElementRef<HTMLDialogElement>;
  @ViewChild('myResultDialog') myResultDialog!: ElementRef<HTMLDialogElement>;

  private productApi = inject(ProductApi)
  public products$ = this.productApi.GetAllProducts()
  public products =
    [
      { id: 1, product_code: "1234567891" },
      { id: 2, product_code: "1234567892" },
      { id: 3, product_code: "1234567893" },
      { id: 4, product_code: "1234567894" },
      { id: 5, product_code: "1234567895" },
    ]
  public delProductIndex: number = -1
  public qrString: string | null = null
  public resultMessage = signal<string | null>(null)
  public isLoading = false;

  constructor() {
    this.products$.subscribe(products => {
      this.products = products
    })
  }


  public addProduct(productCode: string) {
    if (productCode == '') return
    let product_code = productCode.replaceAll('-', '')
    if (product_code.length != 30) {
      this.resultMessage.set('รูปแบบ รหัสสินค้า ไม่ตรงตามที่กำหนด')
      this.myResultDialog.nativeElement.showModal()
      return
    }
    //let id: number = this.products.length + 1
    this.isLoading = true
    let product = [{ id: 0, product_code: product_code }]
    this.productApi.CreateProduct(product).subscribe({
      next: (products) => {
        this.isLoading = false
        this.products.push(...products)
        this.resultMessage.set('เพิ่มสินค้าสำเร็จ')
        this.myResultDialog.nativeElement.showModal()
      },
      error: (err) => {
        this.isLoading = false
        this.resultMessage.set('เกิดข้อผิดพลาด')
        this.myResultDialog.nativeElement.showModal()
      }
    })
    //this.products.push(product)
  }

  private deleteProduct(index: number) {
    if (this.delProductIndex == -1) return
    this.productApi.DeleteProduct(this.products[index].id).subscribe({
      next: (success) => {
        this.isLoading = false
        this.delProductIndex = -1
        this.products.splice(index, 1)
        this.resultMessage.set('ลบสินค้าสำเร็จ')
        this.myResultDialog.nativeElement.showModal()
      },
      error: (err) => {
        this.isLoading = false
        this.resultMessage.set('เกิดข้อผิดพลาด')
        this.myResultDialog.nativeElement.showModal()
      }
    })
  }

  public openQR(productCode: string) {
    this.qrString = productCode
    this.myQrDialog.nativeElement.showModal();
  }

  public closeQR() {
    this.myQrDialog.nativeElement.close();
  }

  public openConfirm(index: number) {
    this.delProductIndex = index
    this.myConfirmDialog.nativeElement.showModal();
  }

  public closeConfirm() {
    this.myConfirmDialog.nativeElement.close();
  }

  public confirmDel() {
    this.deleteProduct(this.delProductIndex)
    this.closeConfirm()
  }

  public closeResult() {
    this.myResultDialog.nativeElement.close();
    this.resultMessage.set(null)
  }



}
