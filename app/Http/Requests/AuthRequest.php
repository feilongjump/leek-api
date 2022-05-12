<?php

namespace App\Http\Requests;

class AuthRequest extends FormRequest
{
    /**
     * Get the validation rules that apply to the request.
     *
     * @return array<string, mixed>
     */
    public function rules(): array
    {
        return match ($this->path()) {
            'api/login' => [
                'username' => 'required|string',
                'password' => 'required|min:6',
            ],
            default => [
                //
            ],
        };
    }
}
