<?php

namespace App\Http\Requests;

class ProjectColumnRequest extends FormRequest
{
    /**
     * Get the validation rules that apply to the request.
     *
     * @return array<string, mixed>
     */
    public function rules()
    {
        return match ($this->getMethod()) {
            'POST', 'PATCH' => [
                'name' => 'required|min:2',
            ],
            default => [
                //
            ],
        };
    }
}
