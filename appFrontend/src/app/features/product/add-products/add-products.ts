import { Component, input, output } from '@angular/core';
import { FormsModule } from '@angular/forms';

@Component({
  selector: 'app-add-products',
  imports: [FormsModule],
  templateUrl: './add-products.html',
  styleUrl: './add-products.css',
})
export class AddProducts {
  public readonly addEvent = output<string>()
  public isLoading = input<boolean>(false)
  public productCode: string = ''

  public addProduct() {
    if (this.productCode != '') {
      this.addEvent.emit(this.productCode)
    }
    this.productCode = ''
  }

  onInput(event: Event) {
    const input = event.target as HTMLInputElement;
    let value = input.value.toUpperCase().replace(/[^A-Z0-9-]/g, '');

    value = value
      .replace(/-/g, '')
      .slice(0, 30)
      .replace(/(.{5})/g, '$1-')
      .slice(0, 35)
      .replace(/-$/, '')

    this.productCode = value;
  }
}
